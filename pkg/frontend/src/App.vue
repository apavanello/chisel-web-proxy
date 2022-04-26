<template>
    <div id="App">
    <title>Chisel Web Proxy</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.7.2/css/bulma.min.css">
      <div class="container">
        <div class="columns">
            <div class="column is-4 is-offset-4">
                <div class="box">
                    <h1 class="title">
                        Chisel Web Proxy
                    </h1>
                        <div class="field">
                            <label class="label">Select environment:</label>
                            <div class="control has-icons-left has-icons-right">
                            <select class="input" @input="setSelectedEnv($event)" :disabled="disabled == 1">
                                <option value="-1" disabled selected="selected"> {{ Benvironment }} </option>
                                <option v-for="environment in environments" :key="environment.id" :value="environment.id">
                                        {{ environment }}
                                </option>
                            </select>
                            </div>
                        </div>
                        <div class="field">
                            <label class="label">Hostname</label>
                            <div class="control has-icons-left has-icons-right">
                                <select class="input" @input="setSelectedHost($event)" @change="setSelectedHost($event)" :v-model="host" :disabled="disabled == 1">
                                    <option value="-1" disabled selected="selected"> {{ Bhost }} </option>
                                    <option v-for="host in hosts" :key="host.id" :value="host.id">
                                        {{ host }}
                                    </option>
                                 </select>
                            </div>
                        </div>
                        <div class="field"  hidden>
                            <label class="label">Local Port</label>
                            <div class="control has-icons-left has-icons-right">
                                <input class="input" type="input" v-model="port" placeholder="8090" :disabled="disabled == 1">
                            </div>
                        </div>
                        <div class="field">
                            <div class="control">
                                <div class="columns">
                                    <div class="column">
                                    <button class="button is-primary" :disabled="disabled == 1"  @click="Connect">Connect</button>
                                    </div>
                                    <div class="column">
                                    <button class="button is-danger is-outlined" :disabled="disabled == 0" @click="Disconnect">Disconnect</button>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div>
                        </div>
                    <div>
                        <p white-space: pre-line> {{message}} </p>
                    </div>
                </div>
            </div>
        </div>
      </div>
      <footer>
        <div class="container">
           <div class="content has-text-centered">
            <p>
              <strong>Chisel Web Proxy</strong> by Apavanello.
            </p>
            <p>
              <a href="https://github.com/apavanello/chisel-web-proxy">Github</a>
            </p>
          </div>
        </div>
        </footer>
    </div>
</template>

<script>

import {ConnectRequest, DisconnectRequest,PreLoadRequest,HostsRequest} from './proto/gateway_pb';
import {ProxyServicePromiseClient,PreLoadServicePromiseClient} from './proto/gateway_grpc_web_pb';

export default{
    
    name: 'App',
    data() {
        return {
            hosts: '',
            port: '8090',
            message: '',
            environments: '',
            myValue: '',
            selectedHost: '',
            disabled: "0",
            host: '',
            environment: '',
            Benvironment: '',
            Bhost: '',
        }
    },
  mounted: 
    function() {
        this.Preload();
    },

  methods: {
    async Preload () {
        this.client = new PreLoadServicePromiseClient('http://localhost:9000');
        const request = new PreLoadRequest();
        try {
            const response = await this.client.preLoad(request);
            let data = response.toObject();

            console.log(data.status);

            if (data.status != "Disconnected") {
                console.log(data.connectrequest)

                    this.host = data.connectrequest.host;
                //    this.port = data.connectrequest.localport
                    this.environment = data.connectrequest.environment;
                    this.disabled = "1";
                    this.message = "Already connected to " + data.connectrequest.environment + " - " + data.connectrequest.host;
                
                //TODO
            }
            
            let envList = [];

            for (var i in data.envList) {
                envList.push(data.envList[i]);
                // more statements
            }

            this.environments = envList


        } catch (err) {
            console.error(err.message)
            this.message = err + " -> Check in Console maybe server is Offline";
            this.disabled = "1";
            throw err
            
        }
    },
    setSelectedEnv(value) {
        this.environment = value.target.value;
        this.GetHosts();
    },
        setSelectedHost(value) {
        this.host = value.target.value;
    },
    async Connect () {

        this.client = new ProxyServicePromiseClient('http://localhost:9000');
        const request = new ConnectRequest();
        request.setHost(this.host);

        console.log("host ->" + this.host);
        console.log("env ->" + this.environment);
        console.log("port ->" + this.port);
        request.setLocalport(this.port);
        request.setEnvironment(this.environment);
        try {
            const res = await this.client.connect(request);
            this.message = res.toObject().status;
            if (res.toObject().hasconnected == true) {
                this.disabled = "1";
            }
        } catch (err) {
            console.error(err.message)
            this.message = err;
            throw err
        }
    },
    async Disconnect () {
        this.client = new ProxyServicePromiseClient('http://localhost:9000');
        const request = new DisconnectRequest();
        const res = await this.client.disconnect(request);
        this.message = res;
        this.Bhost = '';
    //    this.port = ''
        this.Benvironment = '';
        this.disabled = 0;
        location.reload();
    },
    async GetHosts () {
        this.client = new PreLoadServicePromiseClient('http://localhost:9000');
        const request = new HostsRequest();
        request.setEnv(this.environment);

        const response = await this.client.getHosts(request);

        let data = response.toObject()
        let hl = [];

        for (var i in data.hostsList) {
            
            hl.push(data.hostsList[i]);
        }
        this.selectedHost = -1;
        this.hosts = hl;
    },
    updateConnected(port) {
        console.log(port)
        this.port = port;
    },
}
}

</script>