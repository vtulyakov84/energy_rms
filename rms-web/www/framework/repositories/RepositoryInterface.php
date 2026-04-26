// src/Repositories/RepositoryInterface.php
<?php

namespace Repositories;

interface RepositoryInterface
{
    public function find(int $id): ?object;
    public function findAll(): array;
    public function findBy(array $criteria): array;
    public function findOneBy(array $criteria): ?object;
}