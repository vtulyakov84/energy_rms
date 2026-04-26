<?php
namespace framework;

class Autoloader {
	public static function register() {
		ini_set('unserialize_callback_func', 'spl_autoload_call');
		spl_autoload_register(array(new self, 'autoload'));
	}
	
	private static function autoload($class_name) {
		// echo $class_name;
		
		$file_name = getenv('DOCUMENT_ROOT')
			.DIRECTORY_SEPARATOR
			.str_replace('\\', DIRECTORY_SEPARATOR, $class_name).'.php';
		
		// echo $file_name;
		
		if (is_file($file_name)) {
                        // echo "Файл найден $file_name<br />";
			require_once $file_name;
		} else {
			// echo 'framework\Autoloader.php, не найден файл: '.$file_name.'<br />';
			// echo getenv('DOCUMENT_ROOT').'<br />';
			// echo $class_name.'<br />';
			// echo 'Class file <b>\''.$file_name.'\'</b> not found.<br />';
		}
	}
}
?>