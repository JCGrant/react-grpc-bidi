import { grpc } from 'grpc-web-client';
import * as React from 'react';
import { PlayerUpdate, StreamPlayerUpdatesRequest } from './protos/game_pb';
import { GameService } from './protos/game_pb_service';

interface IAppState {
  playerUpdates: PlayerUpdate[];
}

class App extends React.Component<{}, IAppState> {
  constructor(props: {}) {
    super(props);
    this.state = {
      playerUpdates: [],
    }
  }

  public componentDidMount() {
    const client = grpc.client(GameService.StreamPlayerUpdates, {
      host: "http://127.0.0.1:9090",
      transport: grpc.WebsocketTransportFactory,
    });
    client.onMessage(this.addPlayerUpdate);

    client.start();

    const request = new StreamPlayerUpdatesRequest();
    client.send(request);
  }

  public render() {
    return (
      <svg width={800} height={800}>
        <g>
          {this.state.playerUpdates.map((pu, i) => <circle key={i} cx={pu.getX()} cy={pu.getY()} r={2} />)}
        </g>
      </svg>
    );
  }

  private addPlayerUpdate = (pu: PlayerUpdate) => {
    this.setState((state) => ({
      playerUpdates: [
        ...state.playerUpdates,
        pu,
      ]
    }));
  }
}

export default App;
