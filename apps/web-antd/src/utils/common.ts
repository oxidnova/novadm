import dayjs from 'dayjs';
import customParseFormat from 'dayjs/plugin/customParseFormat';
import localizedFormat from 'dayjs/plugin/localizedFormat';

dayjs.extend(localizedFormat);
dayjs.extend(customParseFormat);

const formatDateFromTimestamp = (timestamp: number) => {
  return dayjs(timestamp * 1000).format('YYYY-MM-DD HH:mm:ssZ');
};

const formatDateFromRFC3339 = (date?: string) => {
  if (!date) {
    return '';
  }
  return dayjs(date).format('YYYY-MM-DD HH:mm:ssZ');
};

export { formatDateFromRFC3339, formatDateFromTimestamp };
