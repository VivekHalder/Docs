import React, { useEffect, useRef } from "react";
import "../style.css";

type TextAreaProps = {
    value: string,
    onChange: (event: React.FormEvent<HTMLDivElement>) => void;
}

function TextArea({ value, onChange }: TextAreaProps) {
    const editorRef = useRef<HTMLDivElement>(null);

    useEffect(() => {
        const editor = editorRef.current;

        if (editor) {
            const selection = window.getSelection();
            const range = document.createRange();
            range.selectNodeContents(editor);
            range.collapse(false);
            selection?.removeAllRanges();
            selection?.addRange(range);
        }
    }, [value]);

    return (
        <div>
            <div contentEditable="true" onInput={onChange} dangerouslySetInnerHTML={{ __html: value }} className="editor" ref={editorRef}>
            </div>
        </div>
    )
}

export default TextArea