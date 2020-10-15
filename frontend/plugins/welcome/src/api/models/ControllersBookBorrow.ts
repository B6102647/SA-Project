/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ControllersBookBorrow
 */
export interface ControllersBookBorrow {
    /**
     * 
     * @type {string}
     * @memberof ControllersBookBorrow
     */
    added?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersBookBorrow
     */
    book?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersBookBorrow
     */
    purpose?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersBookBorrow
     */
    user?: number;
}

export function ControllersBookBorrowFromJSON(json: any): ControllersBookBorrow {
    return ControllersBookBorrowFromJSONTyped(json, false);
}

export function ControllersBookBorrowFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersBookBorrow {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'added': !exists(json, 'added') ? undefined : json['added'],
        'book': !exists(json, 'book') ? undefined : json['book'],
        'purpose': !exists(json, 'purpose') ? undefined : json['purpose'],
        'user': !exists(json, 'user') ? undefined : json['user'],
    };
}

export function ControllersBookBorrowToJSON(value?: ControllersBookBorrow | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'added': value.added,
        'book': value.book,
        'purpose': value.purpose,
        'user': value.user,
    };
}

