package scanner

import (
    "encoding/json"
    "os"

    "OrsoNetwork/internal/models"
)

const configFile = "scanner_config.json"

func LoadConfig() (models.ScannerConfig, error) {

    data, err := os.ReadFile(configFile)

    if err != nil {
        return models.ScannerConfig{}, err
    }

    var config models.ScannerConfig

    err = json.Unmarshal(
        data,
        &config,
    )

    if err != nil {
        return models.ScannerConfig{}, err
    }

    return config, nil
}

func SaveConfig(
    config models.ScannerConfig,
) error {

    data, err := json.MarshalIndent(
        config,
        "",
        "    ",
    )

    if err != nil {
        return err
    }

    return os.WriteFile(
        configFile,
        data,
        0644,
    )
}