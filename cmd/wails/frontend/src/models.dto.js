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

class ModelDto {

    /**
     * @param asset<Asset>
     */
    static fromObject(asset) {
        let dto = _.extend(new ModelDto(), asset)

        _.unset(dto, 'SoftwareAsset.Asset')
        _.unset(dto, 'HardwareAsset.Asset')

        mapDeep(dto, d => d.Valid ? d.Time : null,
                o => Object.getPrototypeOf(o) === DeletedAt.prototype)

        return dto
    }
}

export {ModelDto}