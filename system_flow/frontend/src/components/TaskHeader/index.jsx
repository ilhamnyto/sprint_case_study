"use client";

import { useState } from "react";

import { Button } from "@/components/ui/button";
import CalendarInput from "@/components/TaskHeader/CalendarInput";
import { formatISO } from "date-fns";

export default function TaskInput() {
  const [date, setDate] = useState();
  const [todo, setTodo] = useState("");
  const submitHandler = (event) => {
    event.preventDefault();
    let postData = {
      title: todo,
      deadline: date ? formatISO(date, "yyyy-MM-dd'T'HH:mm:ss'Z'") : null,
    };

    fetch(`${process.env.NEXT_PUBLIC_API_HOST}/api/v1/tasks/create`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(postData),
    })
      .then((res) => {
        console.log(res.json());
      })
      .catch((err) => {
        console.log(err);
      });

    setTodo("");
    setDate(null);
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
