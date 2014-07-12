<html>
Name: <?=$_POST["name"]?><br/>

<?
	$parsed_url = parse_url($_POST["name"]);
	parse_str($parsed_url["query"],$output);

	$subscribe_email = array('$addToSet' => array("email" => $_POST["email"]));

	$doc = array("_id" => $output["Item"]);

	echo "<br/>Here 01?<br/>";
	$connection = new MongoClient($_ENV["OPENSHIFT_MONGODB_DB_URL"]);
	var_dump($connection);
	echo "<br/>Here 02?<br/>";
	$db = $connection->neweggtracker;
	var_dump($db);
	echo "<br/>Here 03?<br/>";
	$collection = $db->subscription;
	echo "<br/>Here 04?<br/>";
	var_dump($collection);
	echo "<br/>Here 05?<br/>";
	var_dump($collection->update($doc,$subscribe_email),array('upsert' => true));
	echo "<br/>Here 06?<br/>";
?>
singleFinalPrice
<br/><br/>
</html>
