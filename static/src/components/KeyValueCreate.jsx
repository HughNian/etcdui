import React from 'react'
import { Input, Button } from 'antd'
import { Box } from 'react-polymer-layout'
import { KVPost, KVDelete } from './request'
import { DeleteButton } from './utils'

const KeyValueCreate = React.createClass({
    _createDone(result) {
        this.setState({ key: "", value: "" })
        this.props.update()
    },

    _createKey(e) {
        KVPost(this.props.fullKey(this.state.key), this.state.value, this._createDone)
    },

    _createDir(e) {
        KVPost(this.props.fullKey(this.state.key) + "?dir", null, this._createDone)
    },

    _deleteDone(result) {
        this.props.back()
    },

    _deleteDir() {
        KVDelete(this.state.dir, this._deleteDone)
    },

    getInitialState() {
        return { dir: "", key: "", value: "" }
    },

    _updateDir(props) {
        this.setState({ dir: props.dir, key: "", value: "" })
    },

    componentDidMount() {
        this._updateDir(this.props)
    },

    componentWillReceiveProps(nextProps) {
        if (this.props.dir !== nextProps.dir) {
            this._updateDir(nextProps)
        }
    },

    render() {
        let cantClick = this.state.key === ""
        return (
            <Box vertical style={{ padding: "10px 7px 0px 7px" }}>
                <Box stretch style={{ height: 40 }} >
                    <Input size="large" addonBefore={this.state.dir} placeholder="目录 / 键名" value={this.state.key} onChange={e => this.setState({ key: e.target.value }) } style={{ height: 40 }}/>
                </Box>
                <div style={{ width: "100%", paddingTop: 10, height: 180 }}>
                    <Input type="textarea" rows={4} value={this.state.value} onChange={e => this.setState({ value: e.target.value }) } style={{ height: 160 }}/>
                </div>
                <Box justified >
                    {
                        <Box>
                            <div className="kv-create-button" ><Button type="primary" ghosts size="large" onClick={this._createKey} disabled={cantClick} >创建键</Button></div>
                            <div className="kv-create-button" ><Button type="primary" ghost size="large" onClick={this._createDir} disabled={cantClick}>创建目录</Button></div>
                        </Box>
                    }
                    {
                        this.state.dir === "/" ? null :
                            (<div className="kv-create-button" style={{ paddingRight: 0 }}><DeleteButton name="删除目录" delete={this._deleteDir} /></div>)
                    }
                </Box>
            </Box>
        )
    }
})

module.exports = KeyValueCreate