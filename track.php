<html>
Name: <?=$_POST["name"]?><br/>

<?
	$parsed_url = parse_url($_POST["name"]);
	parse_str($parsed_url["query"],$output);

	$subscribe_email = array('$addToSet' => array("email" => $_POST["email"]));

	$doc = array("_id" => $output["Item"]);

	$connection = new MongoClient();
	var_dump($connection);
	$db = $connection->neweggtracker;
	var_dump($db);
	$collection = $db->subscription;
	var_dump($collection);
	var_dump($collection->update($doc,$subscribe_email));
?>
singleFinalPrice
<br/><br/>
</html>
