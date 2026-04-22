<?php

    require_once './framework/Autoloader.php';
    framework\Autoloader::register();

    use framework\core\PostgreSQL;
    use framework\repositories\UserRepository;

    // Загрузка конфигурации
    $config = require './framework/config/pg.php';

    // Подключение к БД
    $pdo = PostgreSQL::getInstance($config);

    // Создание репозитория
    $userRepository = new UserRepository($pdo);

// Примеры использования

// 1. Найти пользователя по ID
$user = $userRepository->find(1);
if ($user) {
    echo "Пользователь: {$user->getName()}, Email: {$user->getEmail()}\n";
}

// 2. Найти всех пользователей
$allUsers = $userRepository->findAll();
echo "Всего пользователей: " . count($allUsers) . "\n";

// 3. Найти по критериям
$youngUsers = $userRepository->findBy(['age' => 25]);
foreach ($youngUsers as $youngUser) {
    echo "Молодой пользователь: {$youngUser->getName()}\n";
}

// 4. Найти старше 30 лет
$olderUsers = $userRepository->findOlderThan(30);
foreach ($olderUsers as $olderUser) {
    echo "Пользователь старше 30: {$olderUser->getName()}, Возраст: {$olderUser->getAge()}\n";
}

// 5. Поиск по имени
$searchResults = $userRepository->searchByName('John');
foreach ($searchResults as $result) {
    echo "Найден: {$result->getName()}\n";
}

// 6. Пагинация
$page = 1;
$perPage = 10;
$offset = ($page - 1) * $perPage;
$paginatedUsers = $userRepository->findWithPagination($perPage, $offset);
$totalUsers = $userRepository->count();

echo "Страница {$page} из " . ceil($totalUsers / $perPage) . "\n";

// 7. Поиск одного пользователя
$admin = $userRepository->findOneBy(['email' => 'admin@example.com']);
if ($admin) {
    echo "Найден администратор: {$admin->getName()}\n";
}