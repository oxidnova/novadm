import dayjs from 'dayjs';
import customParseFormat from 'dayjs/plugin/customParseFormat';
import localizedFormat from 'dayjs/plugin/localizedFormat';

dayjs.extend(localizedFormat);
dayjs.extend(customParseFormat);

const formatDateFromTimestamp = (timestamp: number) => {
  return dayjs(timestamp * 1000).format('YYYY-MM-DD HH:mm:ssZ');
};

export { formatDateFromTimestamp };
