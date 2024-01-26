import { clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs) {
  return twMerge(clsx(inputs));
}

export async function getData() {
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_HOST}/api/v1/tasks`);

  if (!res.ok) {
    throw new Error("Failed to fetch data");
  }

  return res.json();
}

export async function deleteTask(taskId, isSubTask) {
  let url = isSubTask
    ? `${process.env.NEXT_PUBLIC_API_HOST}/api/v1/subtask/${taskId}`
    : `${process.env.NEXT_PUBLIC_API_HOST}/api/v1/tasks/${taskId}`;

  const res = await fetch(url, {
    method: "DELETE",
  });

  if (!res.ok) {
    throw new Error("Failed to fetch data");
  }

  return res.json();
}

export async function completeTask(taskId, isSubTask) {
  let url = isSubTask
    ? `${process.env.NEXT_PUBLIC_API_HOST}/api/v1/subtask/complete`
    : `${process.env.NEXT_PUBLIC_API_HOST}/api/v1/tasks/complete`;

  const res = await fetch(url, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ id: taskId }),
  });

  if (!res.ok) {
    throw new Error("Failed to fetch data");
  }

  return res.json();
}

export async function addTask({ title, deadline }) {
  let url = `${process.env.NEXT_PUBLIC_API_HOST}/api/v1/tasks/create`;

  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ title: title, deadline: deadline }),
  });

  if (!res.ok) {
    throw new Error("Failed to fetch data");
  }

  return res.json();
}

export async function addSubTask(taskId, title, deadline) {
  let url = `${process.env.NEXT_PUBLIC_API_HOST}/api/v1/subtask/create`;

  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ task_id: taskId, title: title, deadline: deadline }),
  });

  if (!res.ok) {
    throw new Error("Failed to fetch data");
  }

  return res.json();
}

export async function updateTaskOrSubtTask(taskId, title, deadline, isSubTask) {
  let url = isSubTask
    ? `${process.env.NEXT_PUBLIC_API_HOST}/api/v1/subtask/update`
    : `${process.env.NEXT_PUBLIC_API_HOST}/api/v1/tasks/update`;

  const res = await fetch(url, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ id: taskId, title: title, deadline: deadline }),
  });

  if (!res.ok) {
    throw new Error("Failed to fetch data");
  }

  return res.json();
}
