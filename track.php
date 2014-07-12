<html>
Name: <?=$_POST["name"]?><br/>

<?
	$parsed_url = parse_url($_POST["name"]);
	var_dump($parsed_url);
	parse_str($parsed_url["param"],$output);
	var_dump($output);
?>

<br/><br/>
</html>
