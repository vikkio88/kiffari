import { format, formatRelative } from "date-fns";
import { enGB } from "date-fns/locale";

/**
 * @param {Date} date
 * @returns {String} 
 */
export function formatDTL(date) {
    return format(date, "yyyy-MM-dd'T'HH:mm");
}

/**
 * @param {Date} date
 * @returns {String} 
 */
export function formatYHM(date) {
    return format(date, "dd/MM/yyyy HH:mm");
}

/**
 * 
 * @param {Date} date 
 * @returns {String}
 */
export function formatRelativeNow(date) {
    return formatRelative(date, new Date(), { locale: enGB });
}