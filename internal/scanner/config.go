package scanner

import "OrsoNetwork/internal/models"


func (s *Scanner) UpdateConfig(
    config models.ScannerConfig,
) {
    s.Config = config
}