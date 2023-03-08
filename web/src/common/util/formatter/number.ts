import numeral from 'numeral';

export function FormatShortenNumber(number: number) {
  const format = number ? numeral(number).format('0.00a') : '';

  return result(format, '.00');
}

function result(format: any, key = '.00') {
  const isInteger = format.includes(key);

  return isInteger ? format.replace(key, '') : format;
}