<html>
Name: <?=$_POST["name"]?>

<?
	//$_POST["name"]
	//http://www.newegg.com/Product/Product.aspx?Item=N82E16819116989

	phpinfo();

	parse_str($_POST["name"],$output);
	var_dump($output);
?>

<br/><br/>
</html>
