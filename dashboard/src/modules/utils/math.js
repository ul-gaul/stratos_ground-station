/**
 * Rounds a number to a specific decimal precision.
 *
 * @example
 * round(765.4321) // => 765
 * round(765.4321, 1) // => 765.4
 * round(765.4321, 2) // => 765.43
 * round(765.4321, -1) // => 770
 * round(765.4321, -2) // => 800
 *
 * @param {number} value - The value to round
 * @param {number} [precision=0] - Decimal precision
 * @return {number} Rounded value
 */
export function round(value, precision) {
  const e = 10 ** (precision || 0)
  return Math.round(value * e) / e
}