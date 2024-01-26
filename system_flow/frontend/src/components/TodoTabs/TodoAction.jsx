import {
  CheckIcon,
  TrashIcon,
  PlusIcon,
  Pencil1Icon,
} from "@radix-ui/react-icons";

export default function TodoAction({ data, isSubTask }) {
  if (data.completed_at) {
    return (
      <div
        className="bg-red-300 p-1 rounded hover:bg-red-400 transition-all"
        onClick={async () => await deleteTask(data.id, isSubTask)}
      >
        <TrashIcon className="text-white" />
      </div>
    );
  }

  return (
    <div className="flex gap-2 items-center">
      <div
        className="bg-green-300 p-1 rounded hover:bg-green-500 transition-all"
        onClick={async () => await completeTask(data.id, isSubTask)}
      >
        <CheckIcon className="text-white" />
      </div>
      {!isSubTask && (
        <div className="bg-slate-300 p-1 rounded hover:bg-gray-700 transition-all">
          <PlusIcon className="text-white" />
        </div>
      )}
      <div className="bg-slate-300 p-1 rounded hover:bg-gray-700 transition-all">
        <Pencil1Icon className="text-white" />
      </div>
      <div
        className="bg-red-300 p-1 rounded hover:bg-red-400 transition-all"
        onClick={async () => await deleteTask(data.id, isSubTask)}
      >
        <TrashIcon className="text-white" />
      </div>
    </div>
  );
}

async function deleteTask(taskId, isSubTask) {
  let url = isSubTask
    ? `${process.env.NEXT_PUBLIC_API_HOST}/api/v1/subtask/${taskId}`
    : `${process.env.NEXT_PUBLIC_API_HOST}/api/v1/tasks/${taskId}`;

  const res = await fetch(url, {
    method: "DELETE",
  });

  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    throw new Error("Failed to fetch data");
  }

  return res.json();
}

async function completeTask(taskId, isSubTask) {
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
    // This will activate the closest `error.js` Error Boundary
    throw new Error("Failed to fetch data");
  }

  return res.json();
}
