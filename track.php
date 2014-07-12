<html>
<?php

//$_POST["name"]
//http://www.newegg.com/Product/Product.aspx?Item=N82E16819116989

parse_str($_POST["name"],$output);
var_dump($output);

$doc = array(
	"_id" => 
    "name" => "MongoDB",
    "type" => "database",
    "count" => 1,
    "info" => (object)array( "x" => 203, "y" => 102),
    "versions" => array("0.9.7", "0.9.8", "0.9.9")
);

foreach ($_POST as $name => $val)
{
     echo htmlspecialchars($name . ': ' . $val) . "\n";
}

?>

<br/><br/>
Name: <?=$_POST["name"]?>
</html>