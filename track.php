<html>
Name: <?=$_POST["name"]?>

<?
	parse_url($_POST["name"],$output);
	var_dump($output);
?>

<br/><br/>
</html>
