"use client";

import { useState } from "react";
import { CalendarIcon } from "@radix-ui/react-icons";
import { format, formatISO } from "date-fns";

import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";

export default function TaskInput() {
  const [date, setDate] = useState();
  const [todo, setTodo] = useState("");
  const submitHandler = (event) => {
    event.preventDefault();
    if (date) {
      console.log(formatISO(date));
    }
    console.log(todo, date);
  };
  const changeHandler = (event) => {
    setTodo(event.target.value);
  };
  return (
    <div>
      <form className="space-y-2" onSubmit={submitHandler}>
        <input
          type="text"
          className="w-full border border-gray-200 px-3 py-2 rounded-md text-sm"
          placeholder="Task todo..."
          onChange={changeHandler}
          value={todo}
        />
        <CalendarInput date={date} setDate={setDate} />
        <div className="flex justify-end">
          <Button>Add Task</Button>
        </div>
      </form>
    </div>
  );
}

export function CalendarInput({ date, setDate }) {
  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button
          variant={"outline"}
          className={cn(
            "w-full justify-start text-left font-normal",
            !date && "text-muted-foreground"
          )}
        >
          <CalendarIcon className="mr-2 h-4 w-4" />
          {date ? format(date, "PPP") : <span>Pick a deadline</span>}
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-full p-0" align="center">
        <Calendar
          mode="single"
          selected={date}
          onSelect={setDate}
          initialFocus
        />
      </PopoverContent>
    </Popover>
  );
}
