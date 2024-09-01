import React, { useState } from "react";

function ToolBar() {
    const [name, setName] = useState<string>("Untitled Document");

    const handleNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setName(event.target.value)
    }

  return (
    <div className="flex flex-col mb-7">
        <div className="bg-white">
            <input type="text" name="name" value={name} onChange={handleNameChange} className="text-xl p-2 pb-0" />
        </div>
        <div className="flex pb-1">
            <div className="p-2">
                File
            </div>
            <div className="p-2">
                Edit
            </div>
            <div className="p-2">
                View
            </div>
            <div className="p-2">
                Insert
            </div>
            <div className="p-2">
                Format
            </div>
            <div className="p-2">
                Tools
            </div>
            <div className="p-2">
                Extension
            </div>
            <div className="p-2">
                Help
            </div>
        </div>
        <div className="rounded-xl bg-slate-300 h-[35px] mx-1">

        </div>
    </div>
  )
}

export default ToolBar