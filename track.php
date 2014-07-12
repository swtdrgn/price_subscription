<html>
Name: <?=$_POST["name"]?><br/>

<?
	$output = parse_url($_POST["name"]);
	var_dump($output);
?>

<br/><br/>
</html>
