transform "script" do |item|
  script_block = item["arguments"]&.find { |arg| arg["key"] == "scriptBlock" }

  if script_block && script_block["value"]["isLiteral"]
    script_command = script_block["value"]["value"]

    script_command = script_command.sub(/^sh /, "").sub(/^bash /, "")


    {
      name: "set execute permission",
      run: script_command
    }
  else
    nil
  end
end
