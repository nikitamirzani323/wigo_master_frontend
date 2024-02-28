<script>
    import Home from "./Home.svelte";
   
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
    let listBranch = [];
    let record = "";
    let record_message = "";
    let totalrecord = 0;

    async function initapp() {
        const res = await fetch("/api/valid", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "CURR-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
        } else {
            akses_page = true;
            initHome();
        }
    }
    async function initHome() {
        const res = await fetch("/api/warehouse", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            let record_listbranch = json.listbranch;
            record_message = json.message;
            if (record != null) {
                totalrecord = record.length;
                let no = 0
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listHome = [
                        ...listHome,
                        {
                            home_no: no,
                            home_id: record[i]["warehouse_id"],
                            home_idbranch: record[i]["warehouse_idbranch"],
                            home_nmbranch: record[i]["warehouse_nmbranch"],
                            home_name: record[i]["warehouse_name"],
                            home_alamat: record[i]["warehouse_alamat"],
                            home_phone1: record[i]["warehouse_phone1"],
                            home_phone2: record[i]["warehouse_phone2"],
                            home_status: record[i]["warehouse_status"],
                            home_status_css: record[i]["warehouse_status_css"],
                            home_create: record[i]["warehouse_create"],
                            home_update: record[i]["warehouse_update"],
                        },
                    ];
                }
            }
            for (var i = 0; i < record_listbranch.length; i++) {
                listBranch = [
                    ...listBranch,
                    {
                        branch_id: record_listbranch[i]["branch_id"],
                        branch_name: record_listbranch[i]["branch_name"],
                    },
                ];
            }
        } else {
            logout();
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleRefreshData = (e) => {
        listHome = [];
        listBranch = [];
        totalrecord = 0;
        setTimeout(function () {
            initHome();
        }, 500);
    };
    initapp()
</script>

{#if akses_page == true}
<Home
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listHome}
    {listBranch}
    {totalrecord}/>
{/if}