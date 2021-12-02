import {DeletedAt} from "./models";

function mapDeep(object, callback, predicate = null) {
    Object.keys(object).forEach((k) => {
        if (_.isObjectLike(object[k])) {
            if (predicate === null || predicate(object[k])) {
                object[k] = callback(object[k])
                return
            }

            mapDeep(object[k], callback, predicate)
        }
    });
}

class AssetDto {

    /**
     * @param asset<Asset>
     */
    static fromObject(asset) {
        let dto = $.extend(true, new AssetDto(), asset)

        _.unset(dto, 'SoftwareAsset.Asset')
        _.unset(dto, 'HardwareAsset.Asset')

        mapDeep(dto, d => d.Valid ? d.Time : null,
                o => Object.getPrototypeOf(o) === DeletedAt.prototype)

        return dto
    }
}

export {AssetDto}