import React from 'react';
import { Activity } from 'lucide-react';

const LogPanel = ({ logs }) => (
    <div className="bg-slate-800 rounded-xl border border-slate-700 h-[400px] flex flex-col shadow-lg">
        <div className="p-4 border-b border-slate-700 flex justify-between items-center">
            <h3 className="font-semibold text-slate-200 flex items-center gap-2">
                <Activity className="w-4 h-4 text-blue-400" />
                Live Execution Logs
            </h3>
            <span className="text-xs text-slate-500 font-mono">Real-time</span>
        </div>
        <div className="flex-1 overflow-y-auto p-4 font-mono text-sm space-y-2">
            {logs.map((log, i) => (
                <div key={i} className="text-slate-300 border-l-2 border-slate-600 pl-3 py-1">
                    <span className="text-slate-500 mr-2">[{log.time}]</span>
                    {log.msg}
                </div>
            ))}
            <div className="animate-pulse text-blue-400">_</div>
        </div>
    </div>
);

export default LogPanel;
