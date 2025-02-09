import { Task } from "../../utils/consts";
import { FC } from "react";
import TaskTimeStamp from "./components/TaskTimeStamp";
import TaskHostBeacon from "./components/TaskHostBeacon";
import TaskParameters from "./components/TaskParameters";
import TaskResults from "./components/TaskResults";
import UserImageAndName from "../../components/UserImageAndName";
import TaskStatusBadge from "../../components/TaskStatusBadge";
import TaskShells from "./components/TaskShells";

interface TaskCardType {
    task: Task
}

const TaskCard: FC<TaskCardType> = (
    { task }
) => {
    return (
        <div className=" border-2 border-gray-200 px-8 py-5 w-full rounded-lg gap-4 grid grid-cols-1 lg:grid-cols-2">
            <div className="flex flex-col gap-6 col-span-1">
                <div className="flex flex-col gap-1">
                    <div className="flex flex-row gap-2 items-center">
                        <h3 className="text-lg font-semibold">
                            {task.quest?.name}
                        </h3>
                        <TaskStatusBadge task={task} />
                    </div>
                </div>
                <TaskHostBeacon beaconData={task.beacon} />
                <TaskTimeStamp {...task} />
                <TaskParameters quest={task?.quest} />
                <UserImageAndName userData={task?.quest?.creator} />

            </div>
            <div className="flex flex-col gap-6 col-span-1">
                <TaskResults output={task?.output} error={task?.error} quest={task?.quest} />
                <TaskShells shells={task?.shells} />
            </div>
        </div>
    );
};
export default TaskCard;
