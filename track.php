<html>
Name: <?=$_POST["name"]?><br/>

<?
	$parsed_url = parse_url($_POST["name"]);
	parse_str($parsed_url["query"],$output);

	if (strpos(get_headers("http://www.newegg.com/Product/MappingPrice.aspx?Item=" . $output["Item"], 1)["Location"], "http://www.newegg.com/Common/MessagePage.aspx") !== false) {
		echo '["Item is not found."]';
	}
	else {
		try {
			$doc = array("_id" => $output["Item"]);
			$subscribe_email = array('$addToSet' => array("email" => $_POST["email"]));
			$update_options = array('upsert' => true);

			$connection = new MongoClient("mongodb://simon:simon@ds041218.mongolab.com:41218/neweggtracker");
			$db = $connection->neweggtracker;
			$subscription_collection = $db->subscriptions;
			$test = $subscription_collection->update($doc,$subscribe_email,$update_options);
			//echo $connection . "<br/>\n";
			//echo $db . "<br/>\n";
			//echo $subscription_collection  . "<br/>\n";
			//var_dump($test);
			//echo "<br/>\n";
			echo "[]";
			//var_dump($db->lastError());
		} catch (Exception $e) {
			echo 'Caught exception: ',  $e->getMessage(), "\n";
		}
	}

	//$subscribe_email = array('$set' => array('prices' => array('date' => )));
	//$subscribe_email = array('prices' => array('$set' => array('date' => date("y.m.d"), 'price' => 35.23)));
	//$price_collection = $db->prices;
	//$price_collection->update($doc,$subscribe_email,$update_options);
?>
<br/><br/>
</html>
