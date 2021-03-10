
/**
 * @param {number} [ms=0] - Temps d'attente en millisecondes (ms)
 * @returns {Promise<void>}
 */
export const wait = (ms=0) => new Promise((resolve) => setTimeout(resolve, ms));

export default wait;