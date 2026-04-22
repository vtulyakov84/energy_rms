<?php

namespace framework\core;

use PDO;
use PDOException;

class PostgreSQL
{
    private static ?PDO $instance = null;
    private array $config;

    private function __construct(array $config)
    {
        $this->config = $config;
    }

    public static function getInstance(array $config): PDO
    {
        if (self::$instance === null) {
            try {
                $dsn = sprintf(
                    'pgsql:host=%s;port=%s;dbname=%s',
                    $config['host'],
                    $config['port'],
                    $config['dbname']
                );

                self::$instance = new PDO(
                    $dsn,
                    $config['user'],
                    $config['password'],
                    [
                        PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION,
                        PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC,
                        PDO::ATTR_EMULATE_PREPARES => false
                    ]
                );
            } catch (PDOException $e) {
                throw new PDOException("Connection failed: " . $e->getMessage());
            }
        }

        return self::$instance;
    }
}