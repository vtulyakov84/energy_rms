// src/Repositories/UserRepository.php
<?php

namespace Repositories;

use Models\User;
use PDO;
use DateTime;

class UserRepository implements RepositoryInterface
{
    private PDO $connection;
    private string $table = 'users';

    public function __construct(PDO $connection)
    {
        $this->connection = $connection;
    }

    /**
     * Найти пользователя по ID
     */
    public function find(int $id): ?User
    {
        $sql = "SELECT * FROM {$this->table} WHERE id = :id";
        $stmt = $this->connection->prepare($sql);
        $stmt->execute(['id' => $id]);
        
        $result = $stmt->fetch();
        
        return $result ? User::fromArray($result) : null;
    }

    /**
     * Найти всех пользователей
     */
    public function findAll(): array
    {
        $sql = "SELECT * FROM {$this->table} ORDER BY id";
        $stmt = $this->connection->query($sql);
        
        $users = [];
        while ($row = $stmt->fetch()) {
            $users[] = User::fromArray($row);
        }
        
        return $users;
    }

    /**
     * Найти пользователей по критериям
     */
    public function findBy(array $criteria): array
    {
        $conditions = [];
        $params = [];
        
        foreach ($criteria as $key => $value) {
            $conditions[] = "{$key} = :{$key}";
            $params[$key] = $value;
        }
        
        $whereClause = implode(' AND ', $conditions);
        $sql = "SELECT * FROM {$this->table} WHERE {$whereClause}";
        
        $stmt = $this->connection->prepare($sql);
        $stmt->execute($params);
        
        $users = [];
        while ($row = $stmt->fetch()) {
            $users[] = User::fromArray($row);
        }
        
        return $users;
    }

    /**
     * Найти одного пользователя по критериям
     */
    public function findOneBy(array $criteria): ?User
    {
        $users = $this->findBy($criteria);
        return $users[0] ?? null;
    }

    /**
     * Найти пользователей старше определенного возраста
     */
    public function findOlderThan(int $age): array
    {
        $sql = "SELECT * FROM {$this->table} WHERE age > :age ORDER BY age";
        $stmt = $this->connection->prepare($sql);
        $stmt->execute(['age' => $age]);
        
        $users = [];
        while ($row = $stmt->fetch()) {
            $users[] = User::fromArray($row);
        }
        
        return $users;
    }

    /**
     * Поиск по части имени (LIKE)
     */
    public function searchByName(string $searchTerm): array
    {
        $sql = "SELECT * FROM {$this->table} WHERE name ILIKE :search ORDER BY name";
        $stmt = $this->connection->prepare($sql);
        $stmt->execute(['search' => "%{$searchTerm}%"]);
        
        $users = [];
        while ($row = $stmt->fetch()) {
            $users[] = User::fromArray($row);
        }
        
        return $users;
    }

    /**
     * Получить пользователей с пагинацией
     */
    public function findWithPagination(int $limit, int $offset): array
    {
        $sql = "SELECT * FROM {$this->table} ORDER BY id LIMIT :limit OFFSET :offset";
        $stmt = $this->connection->prepare($sql);
        $stmt->bindValue(':limit', $limit, PDO::PARAM_INT);
        $stmt->bindValue(':offset', $offset, PDO::PARAM_INT);
        $stmt->execute();
        
        $users = [];
        while ($row = $stmt->fetch()) {
            $users[] = User::fromArray($row);
        }
        
        return $users;
    }

    /**
     * Получить общее количество пользователей
     */
    public function count(): int
    {
        $sql = "SELECT COUNT(*) as total FROM {$this->table}";
        $stmt = $this->connection->query($sql);
        $result = $stmt->fetch();
        
        return (int)$result['total'];
    }
}