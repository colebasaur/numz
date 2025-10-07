CREATE TABLE IF NOT EXISTS pack_numbers (
  pack_id INTEGER NOT NULL,
  number_id INTEGER NOT NULL,
  rarity INTEGER NOT NULL,
  UNIQUE(pack_id, number_id)
);

CREATE TABLE IF NOT EXISTS collection (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  number_id INTEGER NOT NULL,
  quantity INTEGER NOT NULL,
  UNIQUE(number_id)
);

CREATE INDEX IF NOT EXISTS idx_pack_numbers_pack_id ON pack_numbers(pack_id);
CREATE INDEX IF NOT EXISTS idx_pack_numbers_number_id ON pack_numbers(number_id);
CREATE INDEX IF NOT EXISTS idx_pack_numbers_rarity ON pack_numbers(rarity);

CREATE INDEX IF NOT EXISTS idx_collection_number_id ON pack_numbers(number_id);
