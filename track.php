<html>
Name: <?=$_POST["name"]?><br/>

<?
	$parsed_url = parse_url($_POST["name"]);
	parse_str($parsed_url["query"],$output);
	var_dump($output);

	$subscribe_email = array('$addToSet' => array("email" => $_POST["email"]));

	$doc = array("_id" => $output["Item"]);

	$connection = new MongoClient();
	$db = $connection->neweggtracker;
	$db->update($doc,$subscribe_email);
?>
singleFinalPrice
<br/><br/>
</html>
