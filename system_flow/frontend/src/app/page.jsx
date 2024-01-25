import TaskInput from "@/components/TaskInput";
import TodoTabs from "@/components/TodoTabs";

export default function Home() {
  return (
    <main className="grid place-items-center h-screen bg-slate-100">
      <div className="space-y-2 px-3 w-full max-w-xl">
        <div className="bg-white w-full min-h-[15vh] p-6 shadow rounded">
          <TaskInput />
        </div>
        <div className="">
          <TodoTabs />
        </div>
      </div>
    </main>
  );
}
