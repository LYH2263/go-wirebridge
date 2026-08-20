async function refreshRoutes() {
  const r = await fetch('/api/routes');
  const j = await r.json();
  document.getElementById('routes').textContent = JSON.stringify(j, null, 2);
}
async function refreshStats() {
  const r = await fetch('/api/stats');
  const j = await r.json();
  document.getElementById('stats').textContent = JSON.stringify(j, null, 2);
}
async function sendFrame() {
  const opcode = Number(document.getElementById('opcode').value);
  const payload = document.getElementById('payload').value.trim();
  const bypass = document.getElementById('bypass').checked;
  const flags = bypass ? 8 : 0;
  const enc = await fetch('/api/encode', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ opcode, flags, payload }),
  });
  const ej = await enc.json();
  if (ej.error) {
    document.getElementById('result').textContent = ej.error;
    return;
  }
  const srv = await fetch('/api/serve?hex=1', {
    method: 'POST',
    headers: { 'Content-Type': 'text/plain' },
    body: ej.hex,
  });
  const sj = await srv.json();
  document.getElementById('result').textContent = JSON.stringify(sj, null, 2);
  refreshStats();
}
document.getElementById('btn-refresh').onclick = () => { refreshRoutes(); refreshStats(); };
document.getElementById('btn-send').onclick = sendFrame;
refreshRoutes();
refreshStats();
