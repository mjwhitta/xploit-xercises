<?php
function spam($email) {
  $email[2] = preg_replace("/\./", " dot ", $email[2]);
  $email[2] = preg_replace("/@/", " AT ", $email[2]);
  return $email[2];
}

function markup($filename, $use_me) {
  $contents = file_get_contents($filename);

  $contents = preg_replace(
    "/(\[email (.*)\])/e",
    "spam(\"\\2\")",
    $contents
  );

  $contents = preg_replace("/\[/", "<", $contents);
  $contents = preg_replace("/\]/", ">", $contents);

  return $contents;
}

print markup($argv[1], $argv[2]);
?>
