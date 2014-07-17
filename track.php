<html>
Name: <?=$_POST["name"]?><br/>

<?
	$parsed_url = parse_url($_POST["name"]);
	parse_str($parsed_url["query"],$output);

	if (!empty($output["Item"])) {
		$doc = array("_id" => $output["Item"]);
		$subscribe_email = array('$addToSet' => array("email" => $_POST["email"]));
		$update_options = array('upsert' => true);

		$connection = new MongoClient($_ENV["OPENSHIFT_MONGODB_DB_URL"]);
		$db = $connection->neweggtracker;
		$subscription_collection = $db->subscriptions;
		$subscription_collection->update($doc,$subscribe_email,$update_options);
	} else {
		echo "empty."
	}

	//$subscribe_email = array('$set' => array('prices' => array('date' => )));
	//$subscribe_email = array('prices' => array('$set' => array('date' => date("y.m.d"), 'price' => 35.23)));
	//$price_collection = $db->prices;
	//$price_collection->update($doc,$subscribe_email,$update_options);
?>
singleFinalPrice
<br/><br/>
</html>
