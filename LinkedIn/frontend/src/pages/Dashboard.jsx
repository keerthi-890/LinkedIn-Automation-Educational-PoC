import React, { useState, useEffect } from 'react';
import { Play, Pause, Square, Users, UserPlus, MessageSquare, Shield } from 'lucide-react';
import StatsCard from '../components/StatsCard';
import LogPanel from '../components/LogPanel';

export default function Dashboard() {
    const [status, setStatus] = useState('IDLE'); // IDLE, RUNNING, PAUSED
    const [stats, setStats] = useState({
        profilesFound: 0,
        requestsSent: 0,
        accepted: 0,
        messagesSent: 0
    });
    const [logs, setLogs] = useState([]);
    const [stealthMode, setStealthMode] = useState(true);

    // Poll for stats and logs
    useEffect(() => {
        const interval = setInterval(async () => {
            if (status === 'RUNNING') {
                setStats(prev => ({
                    ...prev,
                    profilesFound: prev.profilesFound + Math.floor(Math.random() * 2),
                    requestsSent: prev.requestsSent + (Math.random() > 0.7 ? 1 : 0)
                }));
                const newLog = {
                    time: new Date().toLocaleTimeString(),
                    msg: `Scanned profile: User_${Math.floor(Math.random() * 1000)}`
                };
                setLogs(prev => [...prev.slice(-19), newLog]);
            }
        }, 2000);
        return () => clearInterval(interval);
    }, [status]);

    const handleStart = async () => {
        setStatus('RUNNING');
        try {
            await fetch('http://localhost:8080/start', { method: 'POST' });
        } catch (e) { console.error("Backend not reachable"); }
    };

    const handleStop = async () => {
        setStatus('IDLE');
        try {
            await fetch('http://localhost:8080/stop', { method: 'POST' });
        } catch (e) { console.error("Backend not reachable"); }
    };

    return (
        <div className="p-8 max-w-7xl mx-auto space-y-8">
            {/* Header */}
            <div className="flex justify-between items-center">
                <div>
                    <h1 className="text-3xl font-bold bg-gradient-to-r from-blue-400 to-indigo-400 bg-clip-text text-transparent">
                        LinkedIn Automation PoC
                    </h1>
                    <p className="text-slate-400 mt-1">Educational Simulation Environment</p>
                </div>
                <div className="flex items-center gap-4">
                    <div className="flex items-center gap-2 px-4 py-2 bg-slate-800 rounded-lg border border-slate-700">
                        <Shield className={`w-4 h-4 ${stealthMode ? 'text-green-400' : 'text-slate-500'}`} />
                        <span className="text-sm font-medium">Stealth Mode: {stealthMode ? 'Active' : 'Disabled'}</span>
                    </div>
                    <div className={`px-3 py-1 rounded-full text-xs font-bold ${status === 'RUNNING' ? 'bg-green-500/20 text-green-400' : 'bg-slate-700 text-slate-400'}`}>
                        {status}
                    </div>
                </div>
            </div>

            {/* Stats Grid */}
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
                <StatsCard title="Profiles Found" value={stats.profilesFound} icon={Users} color="bg-blue-500" />
                <StatsCard title="Requests Sent" value={stats.requestsSent} icon={UserPlus} color="bg-indigo-500" />
                <StatsCard title="Connections" value={stats.accepted} icon={UserPlus} color="bg-emerald-500" />
                <StatsCard title="Messages Sent" value={stats.messagesSent} icon={MessageSquare} color="bg-purple-500" />
            </div>

            {/* Main Content Split */}
            <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
                {/* Controls */}
                <div className="lg:col-span-1 space-y-6">
                    <div className="bg-slate-800 p-6 rounded-xl border border-slate-700 shadow-lg">
                        <h3 className="font-semibold text-lg mb-6 text-white">Automation Controls</h3>

                        <div className="space-y-4">
                            <button
                                onClick={handleStart}
                                disabled={status === 'RUNNING'}
                                className={`w-full py-4 rounded-lg font-bold flex items-center justify-center gap-2 transition-all ${status === 'RUNNING'
                                        ? 'bg-slate-700 text-slate-500 cursor-not-allowed'
                                        : 'bg-blue-600 hover:bg-blue-500 text-white shadow-lg shadow-blue-500/20'
                                    }`}
                            >
                                <Play className="w-5 h-5" /> Start Campaign
                            </button>

                            <button
                                onClick={() => setStatus('PAUSED')}
                                disabled={status !== 'RUNNING'}
                                className="w-full py-3 rounded-lg font-semibold flex items-center justify-center gap-2 bg-slate-700 text-slate-300 hover:bg-slate-600"
                            >
                                <Pause className="w-5 h-5" /> Pause
                            </button>

                            <button
                                onClick={handleStop}
                                className="w-full py-3 rounded-lg font-semibold flex items-center justify-center gap-2 bg-red-500/10 text-red-400 hover:bg-red-500/20 border border-red-500/20"
                            >
                                <Square className="w-5 h-5" /> Stop Emergency
                            </button>
                        </div>

                        <div className="mt-8">
                            <h4 className="text-sm font-medium text-slate-400 mb-4">Configuration</h4>
                            <div className="space-y-3">
                                <label className="flex items-center justify-between p-3 bg-slate-900/50 rounded-lg cursor-pointer">
                                    <span className="text-sm text-slate-300">Humanize Actions</span>
                                    <input type="checkbox" checked={stealthMode} onChange={(e) => setStealthMode(e.target.checked)} className="rounded border-slate-600 text-blue-500 focus:ring-blue-500 bg-slate-700" />
                                </label>
                                <div className="p-3 bg-slate-900/50 rounded-lg">
                                    <div className="flex justify-between text-sm text-slate-300 mb-2">
                                        <span>Daily Limit</span>
                                        <span>{Math.round((stats.requestsSent / 50) * 100)}%</span>
                                    </div>
                                    <div className="h-2 bg-slate-700 rounded-full overflow-hidden">
                                        <div className="h-full bg-blue-500" style={{ width: `${(stats.requestsSent / 50) * 100}%` }}></div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                {/* Logs */}
                <div className="lg:col-span-2">
                    <LogPanel logs={logs} />
                </div>
            </div>
        </div>
    );
}
