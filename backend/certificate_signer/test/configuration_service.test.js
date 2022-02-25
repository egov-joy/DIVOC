describe('should retrieve all mappings if correct configuration layer passed', () => {
    const etcd3 = require('etcd3');
    const config = require('../config/config');
    const { ICD_MAPPINGS_KEYS } = require('../config/constants');
    var mockConfig = {
        CONFIGURATION_LAYER: 'etcd',
        ETCD_URL: 'etcd:2379'
    }
    jest.mock('../config/config', () => {
        return mockConfig;
    });
    console.log = jest.fn();
    const {ConfigLayer, init} = require('../configuration_service');
    var getFn = {
        string: jest.fn()
                .mockReturnValueOnce(new Promise((resolve, reject) => {
                    resolve({})
                }))
                .mockReturnValueOnce(new Promise((resolve, reject) => {
                    resolve([])
                }))
    }
    var mockWatcher = {
        on: jest.fn((event, callback) => {
            callback('some')
        })
    }
    var mockEtcd3Constructor = {
        get: jest.fn().mockImplementation(() => getFn),
        watch: jest.fn().mockImplementation(() => {
            return {
                key: jest.fn().mockImplementation(() => {
                    return {
                        create: jest.fn()
                            .mockReturnValueOnce(new Promise((resolve, reject) => {
                                resolve(mockWatcher);
                            }))
                            .mockReturnValueOnce(new Promise((resolve, reject) => {
                                resolve(mockWatcher);
                            }))
                    }
                })
            }
        })
    };
    jest.mock('etcd3', () => {
        return {
            Etcd3: jest.fn().mockImplementation(() => mockEtcd3Constructor)
        }
    });

    beforeEach(() => {
        init();
    });
    
    test('should initialise two watchers each for ICD and VACCINE_ICD mapping keys', () => {
        expect(etcd3.Etcd3).toHaveBeenCalled();
        expect(mockEtcd3Constructor.watch).toHaveBeenCalledTimes(2);
    });
    
    test('should fetch values of ICD Mapping and VACCINE_ICD mapping from etcd', async() => {
        const ICD = await (new ConfigLayer()).getICDMappings(ICD_MAPPINGS_KEYS.ICD);
        const VACCINE_ICD = await (new ConfigLayer()).getICDMappings(ICD_MAPPINGS_KEYS.VACCINE_ICD);
        expect(ICD).toEqual({})
        expect(VACCINE_ICD).toEqual([])
        expect(mockEtcd3Constructor.get).toHaveBeenCalledTimes(2);
        expect(mockEtcd3Constructor.get).toHaveBeenCalledWith(ICD_MAPPINGS_KEYS.ICD);
        expect(mockEtcd3Constructor.get).toHaveBeenCalledWith(ICD_MAPPINGS_KEYS.VACCINE_ICD);
    });
});

describe('wrong environment variable for configuration layer', () => {
    const { ICD_MAPPINGS_KEYS } = require('../config/constants');
    const OLD_ENV = process.env;
    jest.resetModules();
    var mockConfig = {
        CONFIGURATION_LAYER: 'etc',
        ETCD_URL: 'etcd:2379'
    }
    jest.mock('../config/config', () => {
        return mockConfig;
    });
    const services = require('../configuration_service');
    beforeEach(() => {
        services.init();
    });
    afterEach(() => {
        process.env = OLD_ENV;
    });

    test('should return null if wrong configuration passed', async() => {
        const mapping = await (new services.ConfigLayer()).getICDMappings(ICD_MAPPINGS_KEYS.ICD);
        expect(mapping).toEqual(null);
    })
});