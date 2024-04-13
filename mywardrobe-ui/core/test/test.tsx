import Axios from 'axios';


export default async function test() {
    Axios.post('/api/test/testauth').then(
        res => {
            console.log(res.data);
        }
    )
}