<script>
    import Home from "./Home.svelte";
   
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
    let listPage = [];
    let listCurr = [];
    let listBranch = [];
    let listDepartement = [];
    let search = "";
    let record = "";
    let record_message = "";
    let totalrecord = 0;
    let totalrecordall = 0;
    let totalpaging = 0;
    let perpage = 0;
    let page = 0;

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
    async function initHome(e) {
        listHome = [];
        listCurr = [];
        listBranch = [];
        const res = await fetch("/api/rfq", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                rfq_search: e,
                rfq_page : parseInt(page)
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            let record_listcurr = json.listcurr;
            let record_listbranch = json.listbranch;
            record_message = json.message;
            perpage = json.perpage;
            totalrecordall = json.totalrecord;
            if (record != null) {
                totalpaging = Math.ceil(parseInt(totalrecordall) / parseInt(perpage))
                totalrecord = record.length;
                let no = 0
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listHome = [
                        ...listHome,
                        {
                            home_no: no,
                            home_id: record[i]["rfq_id"],
                            home_date: record[i]["rfq_date"],
                            home_idbranch: record[i]["rfq_idbranch"],
                            home_idvendor: record[i]["rfq_idvendor"],
                            home_idcurr: record[i]["rfq_idcurr"],
                            home_nmbranch: record[i]["rfq_nmbranch"],
                            home_nmvendor: record[i]["rfq_nmvendor"],
                            home_tipedoc: record[i]["rfq_tipedoc"],
                            home_totalitem: record[i]["rfq_totalitem"],
                            home_totalrfq: record[i]["rfq_totalrfq"],
                            home_status: record[i]["rfq_status"],
                            home_status_css: record[i]["rfq_status_css"],
                            home_create: record[i]["rfq_create"],
                            home_update: record[i]["rfq_update"],
                        },
                    ];
                }
                listPage = [];
                for(var i=1;i<=totalpaging;i++){
                    listPage = [
                        ...listPage,
                        {
                            page_id: i,
                            page_value: ((i*perpage)-perpage),
                            page_display: i + " Of " + perpage*i,
                        },
                    ];
                }
            }
            for (var i = 0; i < record_listcurr.length; i++) {
                listCurr = [
                    ...listCurr,
                    {
                        curr_id: record_listcurr[i]["curr_id"],
                    },
                ];
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
        listCurr = [];
        listBranch = [];
        totalrecord = 0;
        setTimeout(function () {
            initHome();
        }, 500);
    };
    const handleSearch = (e) => {
        search = e.detail.searchHome;
        initHome(search,"")
    };
    const handlePaging = (e) => {
        page = e.detail.page
        initHome("")
    };
    initapp()
</script>

{#if akses_page == true}
<Home
    on:handlePaging={handlePaging}
    on:handleSearch={handleSearch}
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listPage}
    {listHome}
    {listCurr}
    {listBranch}
    {listDepartement}
    {totalrecord}/>
{/if}