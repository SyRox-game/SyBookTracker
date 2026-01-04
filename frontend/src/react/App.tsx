import { useEffect, useState } from 'react';
import { GetAllTracks, AddTrack, DeleteTrack,IsHighestID} from "../../wailsjs/go/main/App";
import { main } from "../../wailsjs/go/models";
import '../css/App.css';

function App() {
  const [trackers, setTrackers] = useState<main.Tracker[]>([]);
  const [showModal, setShowModal] = useState(false);
  const [status, setStatus] = useState<'written' | 'not written'>('written');
  const [wordCount, setWordCount] = useState<string>("");

  useEffect(() => {
    GetAllTracks().then((d) => setTrackers(d));
  }, []);

  const openModal = () => setShowModal(true);
  const closeModal = () => setShowModal(false);

  const handleAddTrack = async () => {
    const today = new Date();
    const formattedDate = `${today.getDate().toString().padStart(2, '0')}.${(today.getMonth()+1).toString().padStart(2,'0')}.${today.getFullYear()}`;
    
    try {
      const word = parseInt(wordCount)
      const newTrack = await AddTrack(formattedDate, status, word);
      setTrackers(prev => [...prev, newTrack]);
      closeModal();
      setWordCount("");
      setStatus('written');
    } catch (err) {
      console.error("Failed to add track:", err);
    }
  };

  const handleDeleteTrack = async (trackID: number) => {
    const isHighest = await IsHighestID(trackID)
    if (!isHighest){
      alert("You can only delete the youngest entry")
      return
    }
    setTrackers(prev => prev.filter((n) => n.id !== trackID));
    await DeleteTrack(trackID);
  };

  return (
    <div id="app">
      <div id="notespace-list">
        {trackers?.map((val) => (
          <div
            className='TrackEntry'
            key={val.id}
            style={{
              backgroundColor:
                val.status === 'written'
                  ? 'rgba(76, 175, 80, 0.2)'
                  : 'rgba(255, 77, 79, 0.2)',
              transition: 'background-color 0.3s ease',
            }}
          >
            <div>{val.date} | {val.status}</div>
            <div>Written Words: {val.writtenwords}</div>
            <div>Collected Words after today: {val.wordstatus}</div>
            <button onClick={() => handleDeleteTrack(val.id)}>DELETE</button>
          </div>
        ))}

      </div>

      <div id='addTrack'>
        <button onClick={openModal}>Add Track</button>
      </div>

      {/* Modal */}
      {showModal && (
        <div className="modal-overlay">
          <div
            className="modal"
            style={{
              backgroundColor:
                status === "written" ? "rgba(76, 175, 80, 0.2)" : "rgba(255, 77, 79, 0.2)",
              color: "#e0e0e0",
            }}
          >
            <h2>Add New Track</h2>
            <label>
              Status:
              <select
                value={status}
                onChange={(e) =>
                  setStatus(e.target.value as "written" | "not written")
                }
              >
                <option value="written">Written</option>
                <option value="not written">Not Written</option>
              </select>
            </label>
        
            <label>
              Word Count:
              <input
                type="number"
                value={wordCount}
                min={0}
                placeholder='Current amount of words'
                onChange={(e) => setWordCount(e.target.value)}
              />
            </label>
        
            <div className="modal-buttons">
              <button onClick={handleAddTrack}>Add</button>
              <button onClick={closeModal} className="cancel-btn">
                Cancel
              </button>
            </div>
          </div>
        </div>

      )}
    </div>
  );
}

export default App;
