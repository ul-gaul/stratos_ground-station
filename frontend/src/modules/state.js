import { reactive } from 'vue'
import { round } from '@/modules/utils/math'
import { Log } from '@wailsapp/runtime'

/**
 * @typedef {Object} TransformedData
 * @property {number} time
 * @property {number} altitude
 * @property {number} pressure
 */

/**/
export const state = reactive({
  raw: [],
  times: [],
  altitudes: [],
  pressures: [],
  // Latitude & longitude skipped (they're not displayed in a line chart)
  internalTemps: [],
  externalTemps: [],
  
  speeds: [], // FIXME What kind/direction of speed? Angular speed? Ascension speed?...
  payloads: [], // FIXME Payload? What's the payload?
})

/**
 * Transforms data received from Go (backend) into chart data
 * and add it to the current data.
 *
 * @param {Array<TransformedData>} data
 */
export function onDataReceived(data) {
  Log.Info('Data received!')
  console.log('data received!')

  const extractKey = (key, dflt = 0) => (obj) => (obj[key] ?? dflt);
  extractKey.rnd = (key, dec = 0) => (obj) => round(extractKey(key)(obj), dec)

  state.raw.push(...data);
  state.times.push(...data.map(extractKey('time')));
  state.altitudes.push(...data.map(extractKey.rnd('altitude', 2)));
  state.pressures.push(...data.map(extractKey.rnd('pressure', 2)));
  state.internalTemps.push(...data.map(extractKey.rnd('internalTemp', 1)));
  state.externalTemps.push(...data.map(extractKey.rnd('externalTemp', 1)));
  
  state.speeds.push(...data.map(extractKey('speeds')));
  state.payloads.push(...data.map(extractKey('payloads')));
}