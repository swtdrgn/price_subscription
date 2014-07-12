<html>
Name: <?=$_POST["name"]?><br/>

<?
	$parsed_url = parse_url($_POST["name"]);
	parse_str($parsed_url,$output);
	var_dump($output);
?>

<br/><br/>
</html>
