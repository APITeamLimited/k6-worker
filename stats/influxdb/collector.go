/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2016 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package influxdb

import (
	"context"
	"errors"
	"sync"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/loadimpact/k6/lib"
	"github.com/loadimpact/k6/stats"
	"github.com/sirupsen/logrus"
)

// Verify that Collector implements lib.Collector
var _ lib.Collector = &Collector***REMOVED******REMOVED***

type Collector struct ***REMOVED***
	Client    client.Client
	Config    Config
	BatchConf client.BatchPointsConfig

	buffer      []stats.Sample
	bufferLock  sync.Mutex
	wg          sync.WaitGroup
	semaphoreCh chan struct***REMOVED******REMOVED***
***REMOVED***

func New(conf Config) (*Collector, error) ***REMOVED***
	cl, err := MakeClient(conf)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	batchConf := MakeBatchConfig(conf)
	if conf.ConcurrentWrites.Int64 <= 0 ***REMOVED***
		return nil, errors.New("influxdb's ConcurrentWrites must be a positive number")
	***REMOVED***
	return &Collector***REMOVED***
		Client:      cl,
		Config:      conf,
		BatchConf:   batchConf,
		semaphoreCh: make(chan struct***REMOVED******REMOVED***, conf.ConcurrentWrites.Int64),
	***REMOVED***, nil
***REMOVED***

func (c *Collector) Init() error ***REMOVED***
	// Try to create the database if it doesn't exist. Failure to do so is USUALLY harmless; it
	// usually means we're either a non-admin user to an existing DB or connecting over UDP.
	_, err := c.Client.Query(client.NewQuery("CREATE DATABASE "+c.BatchConf.Database, "", ""))
	if err != nil ***REMOVED***
		logrus.WithError(err).Debug("InfluxDB: Couldn't create database; most likely harmless")
	***REMOVED***

	return nil
***REMOVED***

func (c *Collector) Run(ctx context.Context) ***REMOVED***
	logrus.Debug("InfluxDB: Running!")
	ticker := time.NewTicker(time.Duration(c.Config.PushInterval.Duration))
	for ***REMOVED***
		select ***REMOVED***
		case <-ticker.C:
			c.wg.Add(1)
			go c.commit()
		case <-ctx.Done():
			c.wg.Add(1)
			go c.commit()
			c.wg.Wait()
			return
		***REMOVED***
	***REMOVED***
***REMOVED***

func (c *Collector) Collect(scs []stats.SampleContainer) ***REMOVED***
	c.bufferLock.Lock()
	defer c.bufferLock.Unlock()
	for _, sc := range scs ***REMOVED***
		c.buffer = append(c.buffer, sc.GetSamples()...)
	***REMOVED***
***REMOVED***

func (c *Collector) Link() string ***REMOVED***
	return c.Config.Addr.String
***REMOVED***

func (c *Collector) commit() ***REMOVED***
	defer c.wg.Done()
	c.bufferLock.Lock()
	samples := c.buffer
	c.buffer = nil
	c.bufferLock.Unlock()
	// let first get the data and then wait our turn
	c.semaphoreCh <- struct***REMOVED******REMOVED******REMOVED******REMOVED***
	defer func() ***REMOVED***
		<-c.semaphoreCh
	***REMOVED***()
	logrus.Debug("InfluxDB: Committing...")
	logrus.WithField("samples", len(samples)).Debug("InfluxDB: Writing...")

	batch, err := c.batchFromSamples(samples)
	if err != nil ***REMOVED***
		return
	***REMOVED***

	logrus.WithField("points", len(batch.Points())).Debug("InfluxDB: Writing...")
	startTime := time.Now()
	if err := c.Client.Write(batch); err != nil ***REMOVED***
		logrus.WithError(err).Error("InfluxDB: Couldn't write stats")
	***REMOVED***
	t := time.Since(startTime)
	logrus.WithField("t", t).Debug("InfluxDB: Batch written!")
***REMOVED***

func (c *Collector) extractTagsToValues(tags map[string]string, values map[string]interface***REMOVED******REMOVED***) map[string]interface***REMOVED******REMOVED*** ***REMOVED***
	for _, tag := range c.Config.TagsAsFields ***REMOVED***
		if val, ok := tags[tag]; ok ***REMOVED***
			values[tag] = val
			delete(tags, tag)
		***REMOVED***
	***REMOVED***
	return values
***REMOVED***

func (c *Collector) batchFromSamples(samples []stats.Sample) (client.BatchPoints, error) ***REMOVED***
	batch, err := client.NewBatchPoints(c.BatchConf)
	if err != nil ***REMOVED***
		logrus.WithError(err).Error("InfluxDB: Couldn't make a batch")
		return nil, err
	***REMOVED***

	type cacheItem struct ***REMOVED***
		tags   map[string]string
		values map[string]interface***REMOVED******REMOVED***
	***REMOVED***
	cache := map[*stats.SampleTags]cacheItem***REMOVED******REMOVED***
	for _, sample := range samples ***REMOVED***
		var tags map[string]string
		var values = make(map[string]interface***REMOVED******REMOVED***)
		if cached, ok := cache[sample.Tags]; ok ***REMOVED***
			tags = cached.tags
			for k, v := range cached.values ***REMOVED***
				values[k] = v
			***REMOVED***
		***REMOVED*** else ***REMOVED***
			tags = sample.Tags.CloneTags()
			c.extractTagsToValues(tags, values)
			cache[sample.Tags] = cacheItem***REMOVED***tags, values***REMOVED***
		***REMOVED***
		values["value"] = sample.Value
		p, err := client.NewPoint(
			sample.Metric.Name,
			tags,
			values,
			sample.Time,
		)
		if err != nil ***REMOVED***
			logrus.WithError(err).Error("InfluxDB: Couldn't make point from sample!")
			return nil, err
		***REMOVED***
		batch.AddPoint(p)
	***REMOVED***

	return batch, err
***REMOVED***

// Format returns a string array of metrics in influx line-protocol
func (c *Collector) Format(samples []stats.Sample) ([]string, error) ***REMOVED***
	var metrics []string
	batch, err := c.batchFromSamples(samples)

	if err != nil ***REMOVED***
		return metrics, err
	***REMOVED***

	for _, point := range batch.Points() ***REMOVED***
		metrics = append(metrics, point.String())
	***REMOVED***

	return metrics, nil
***REMOVED***

// GetRequiredSystemTags returns which sample tags are needed by this collector
func (c *Collector) GetRequiredSystemTags() stats.SystemTagSet ***REMOVED***
	return stats.SystemTagSet(0) // There are no required tags for this collector
***REMOVED***

// SetRunStatus does nothing in the InfluxDB collector
func (c *Collector) SetRunStatus(status lib.RunStatus) ***REMOVED******REMOVED***
