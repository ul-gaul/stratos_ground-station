/**
 * Génère des coordonnées au hasard.
 *
 * @param {object} [opts] - Options
 * @param {number} [opts.qty=10] - Nombre de coordonnées à générer
 * @param {boolean} [opts.ordered=true] - Si les coords sont en ordre de X
 * @param {number} [opts.minX=-100] - Valeur minimale de X
 * @param {number} [opts.maxX=100] - Valeur maximale de X
 * @param {number} [opts.minY=-100] - Valeur minimale de Y
 * @param {number} [opts.maxY=100] - Valeur maximale de Y
 * @return {Array<{x: number, y: number}>} Liste de coordonnées
 */
export function coords(opts = {}) {
  opts.qty = Math.abs(opts.qty ?? 10);
  opts.ordered ??= true;
  opts.minX ??= -100;
  opts.maxX ??= 100;
  opts.minY ??= -100;
  opts.maxY ??= 100;
  
  const minX   = Math.min(opts.minX, opts.maxX),
        maxX   = Math.max(opts.minX, opts.maxX),
        minY   = Math.min(opts.minY, opts.maxY),
        maxY   = Math.max(opts.minY, opts.maxY),
        rangeX = maxX - minX,
        rangeY = maxY - minY,
        coords = [];
  
  for (let i = 0; i < opts.qty; i++) {
    coords.push({
      x: Math.random() * rangeX + minX,
      y: Math.random() * rangeY + minY
    });
  }
  
  if (opts.ordered)
    coords.sort((a, b) => b.x - a.x);
  
  return coords;
} // END coords()


/**
 * Génère une couleur RGBA au hasard.
 * @example
 * rgba() // => 'rgba(10, 233, 108, 0.8)'
 * rgba(1) // => 'rgba(88, 33, 201, 1)'
 *
 * @param {number} [alpha] - Transparence (générée au hasard si non-spécifié)
 * @return {string}
 */
export function rgba(alpha) {
  const rgb = [
    Math.floor(Math.random() * 256),
    Math.floor(Math.random() * 256),
    Math.floor(Math.random() * 256)
  ];
  
  if (typeof alpha !== 'number')
    alpha = Math.round(Math.random() * 100) / 100;
  alpha = Math.min(Math.max(alpha, 0), 1);
  
  return `rgba(${rgb.join(', ')}, ${alpha})`;
} // END rgba()

/**
 * Génère une couleur RGB au hasard.
 * @example
 * rgb() // => 'rgb(43, 78, 199)'
 *
 * @return {string}
 */
export function rgb() {
  const rgb = [
    Math.floor(Math.random() * 256),
    Math.floor(Math.random() * 256),
    Math.floor(Math.random() * 256)
  ];
  
  return `rgb(${rgb.join(', ')})`;
} // END rgb()

export default {
  coords,
  rgba,
  rgb
}