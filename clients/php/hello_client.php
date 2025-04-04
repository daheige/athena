<?php
require dirname(__FILE__) . '/vendor/autoload.php';

const GRPC_ADDRESS = '127.0.0.1:8081';


/* % php hello_client.php
     check App\Grpc\GPBMetadata\Hello\HelloReq exist
     bool(true)
     status code: 0
     name:hello,world
     call ok
*/

function greet($id)
{
    $client = new App\Grpc\Hello\GreeterServiceClient(GRPC_ADDRESS, [
        'credentials' => Grpc\ChannelCredentials::createInsecure(),
    ]);

    echo "check App\Grpc\GPBMetadata\Hello\HelloReq exist" . PHP_EOL;
    var_dump(class_exists("App\Grpc\Hello\HelloReq"));
    $request = new App\Grpc\Hello\HelloReq();
    $request->setId($id);

    list($reply, $status) = $client->SayHello($request)->wait();
    echo 'status code: ' . $status->code;
    if ($status->details) {
        echo ', details: ' . $status->details;
    }

    echo PHP_EOL;
    if ($status->metadata) {
        echo 'Meta data: ' . PHP_EOL;
        print_r($status->metadata);
    }

    echo "name:" . $reply->getName();
    echo PHP_EOL;

    return $reply->getMessage();
}

$name = !empty($argv[1]) ? $argv[1] : 1;
echo greet($name) . PHP_EOL;
