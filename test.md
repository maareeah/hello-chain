# Test
hello post Create message

```
    hellod tx hello create-post title-01 body-01 jan-12 --from alice
```

Response

```
    code: 0
    codespace: ""
    data: 12240A222F68656C6C6F2E68656C6C6F2E4D7367437265617465506F7374526573706F6E7365
    events:
    - attributes:
    - index: true
        key: ZmVl
        value: ""
    - index: true
        key: ZmVlX3BheWVy
        value: Y29zbW9zMTNrenlqcnNycDBuM3UwejNzYXlnc3MwdnBxZ2NqNTA5eDh4ZG4y
    type: tx
    - attributes:
    - index: true
        key: YWNjX3NlcQ==
        value: Y29zbW9zMTNrenlqcnNycDBuM3UwejNzYXlnc3MwdnBxZ2NqNTA5eDh4ZG4yLzE=
    type: tx
    - attributes:
    - index: true
        key: c2lnbmF0dXJl
        value: cDN2MlloSW4wZHo4dnhvNFhlVmxKOGlxSHRaaXJiRmRINmVaS1Z3U2F5WUJqSjVYK0s5aHprclowd1F2a1J2bVJrMFhVTGlGQ0V3U1J2dVNZRkxDdnc9PQ==
    type: tx
    - attributes:
    - index: true
        key: YWN0aW9u
        value: L2hlbGxvLmhlbGxvLk1zZ0NyZWF0ZVBvc3Q=
    type: message
    gas_used: "52244"
    gas_wanted: "200000"
    height: "46778"
    info: ""
    logs:
    - events:
    - attributes:
        - key: action
        value: /hello.hello.MsgCreatePost
        type: message
    log: ""
    msg_index: 0
    raw_log: '[{"msg_index":0,"events":[{"type":"message","attributes":[{"key":"action","value":"/hello.hello.MsgCreatePost"}]}]}]'
    timestamp: ""
    tx: null
    txhash: DA9E4332092EEF2F07CB73070036B783CC86DD0BE8456974CA29109E56BF6E71
```

# check txns in blockchin

http://localhost:26657/tx?hash=0xDA9E4332092EEF2F07CB73070036B783CC86DD0BE8456974CA29109E56BF6E71

```
    {
        "jsonrpc": "2.0",
        "id": -1,
        "result": {
            "hash": "DA9E4332092EEF2F07CB73070036B783CC86DD0BE8456974CA29109E56BF6E71",
            "height": "46778",
            "index": 0,
            "tx_result": {
                "code": 0,
                "data": "EiQKIi9oZWxsby5oZWxsby5Nc2dDcmVhdGVQb3N0UmVzcG9uc2U=",
                "log": "[{\"msg_index\":0,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/hello.hello.MsgCreatePost\"}]}]}]",
                "info": "",
                "gas_wanted": "200000",
                "gas_used": "52244",
                "events": [
                    {
                        "type": "tx",
                        "attributes": [
                            {
                                "key": "ZmVl",
                                "value": null,
                                "index": true
                            },
                            {
                                "key": "ZmVlX3BheWVy",
                                "value": "Y29zbW9zMTNrenlqcnNycDBuM3UwejNzYXlnc3MwdnBxZ2NqNTA5eDh4ZG4y",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "tx",
                        "attributes": [
                            {
                                "key": "YWNjX3NlcQ==",
                                "value": "Y29zbW9zMTNrenlqcnNycDBuM3UwejNzYXlnc3MwdnBxZ2NqNTA5eDh4ZG4yLzE=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "tx",
                        "attributes": [
                            {
                                "key": "c2lnbmF0dXJl",
                                "value": "cDN2MlloSW4wZHo4dnhvNFhlVmxKOGlxSHRaaXJiRmRINmVaS1Z3U2F5WUJqSjVYK0s5aHprclowd1F2a1J2bVJrMFhVTGlGQ0V3U1J2dVNZRkxDdnc9PQ==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "YWN0aW9u",
                                "value": "L2hlbGxvLmhlbGxvLk1zZ0NyZWF0ZVBvc3Q=",
                                "index": true
                            }
                        ]
                    }
                ],
                "codespace": ""
            },
            "tx": "CmoKaAoaL2hlbGxvLmhlbGxvLk1zZ0NyZWF0ZVBvc3QSSgotY29zbW9zMTNrenlqcnNycDBuM3UwejNzYXlnc3MwdnBxZ2NqNTA5eDh4ZG4yEgh0aXRsZS0wMRoHYm9keS0wMSIGamFuLTEyElgKUApGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQJZxaeRh7OaWzklshxRY3yU//q26ZJHBtKXDteH/WFweBIECgIIARgBEgQQwJoMGkCne/ZiEifR3Py/Gjhd5WUnyKoe1mKtsV0fp5kpXBJrJgGMnlf4r2HOStnTBC+RG+ZGTRdQuIUITBJG+5JgUsK/"
        }
    }

```
