-- Create syntax for TABLE 'addresses'
CREATE TABLE `addresses` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `address` varchar(129) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq-address` (`address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='Index address';

-- Create syntax for TABLE 'hashes'
CREATE TABLE `hashes` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `hash` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq-hash` (`hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='Index tx hash';

-- Create syntax for TABLE 'tokens'
CREATE TABLE `tokens` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `symbol` varchar(129) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT 'desc',
  `supply` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `decimals` tinyint(2) unsigned NOT NULL,
  `contract` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT 'aontract address',
  `icon` varchar(255) COLLATE utf8_unicode_ci DEFAULT '' COMMENT 'icon url',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `uniq-contract` (`contract`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='Support tokens. Note: eth.id=1';

-- Create syntax for TABLE 'transactions'
CREATE TABLE `transactions` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `hash_id` bigint(20) unsigned NOT NULL,
  `hash_index` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'tx(common, internal) index in this tx',
  `from_address_id` bigint(20) NOT NULL COMMENT 'actual from address',
  `to_address_id` bigint(20) NOT NULL COMMENT 'actual to address',
  `block_number` bigint(20) unsigned NOT NULL COMMENT 'block number where this transaction was in',
  `amount` varchar(255) NOT NULL DEFAULT '' COMMENT 'transaction amount',
  `token_id` bigint(20) unsigned NOT NULL COMMENT 'token id',
  `gas_used` varchar(255) NOT NULL DEFAULT '0' COMMENT 'tx_receipt.gas_used',
  `gas_price` varchar(255) NOT NULL DEFAULT '0' COMMENT 'tx.gas_price',
  `value` varchar(255) NOT NULL DEFAULT '0' COMMENT 'tx.value',
  `inout_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0: default; 1: in/out(transfer); 2: contract(not transfer)',
  `block_timestamp` int(10) unsigned NOT NULL COMMENT 'the unix timestamp for when the block was collated.',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx-from_address_id` (`from_address_id`),
  KEY `idx-to_address_id` (`to_address_id`),
  KEY `idx-token_id` (`token_id`),
  KEY `idx-hash_id` (`hash_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Transactions for eth and tokens';