/**
 * @module speedboat
 */

/**
 * Runs code in a group.
 * @param  ***REMOVED***string***REMOVED***   name   Name of the group.
 * @param  ***REMOVED***Function***REMOVED*** fn     Group body.
 * @param  ***REMOVED***any***REMOVED***      [cond] If given, the group will be skipped if falsy.
 * @return ***REMOVED***any***REMOVED***             The return value of fn().
 */
export function group(name, fn, cond) ***REMOVED***
	if (cond !== undefined && !cond) ***REMOVED***
		return
	***REMOVED***

	return __jsapi__.DoGroup(name, fn);
***REMOVED***

/**
 * Runs tests on a value.
 * @param  ***REMOVED***any***REMOVED***    val     Value to test.
 * @param  ***REMOVED***...Object***REMOVED*** sets Sets of tests.
 */
export function test(val, ...sets) ***REMOVED***
	return __jsapi__.DoTest(val, ...sets);
***REMOVED***

/**
 * Sleeps for the specified duration.
 * @param  ***REMOVED***Number***REMOVED*** secs Duration, in seconds.
 */
export function sleep(secs) ***REMOVED***
	__jsapi__.Sleep(secs);
***REMOVED***

/**
 * Asserts that a value is truthy.
 * @param  ***REMOVED***any***REMOVED***    exp   Expression result.
 * @param  ***REMOVED***string***REMOVED*** [err] Error message.
 * @throws ***REMOVED***Error***REMOVED***        If exp is falsy.
 */
export function _assert(exp, err = "assertion failed") ***REMOVED***
	if (!exp) ***REMOVED***
		throw new Error(err);
	***REMOVED***
***REMOVED***

export default ***REMOVED***
	group: group,
	test: test,
	sleep: sleep,
***REMOVED***;
