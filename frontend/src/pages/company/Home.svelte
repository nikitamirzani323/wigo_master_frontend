<script> 
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
    import Input_custom from '../../components/InputCustom.svelte' 
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";

    
	export let table_header_font = ""
	export let table_body_font = ""
	export let token = ""
	export let listHome = []
	export let listCurrency = []
    export let listPage = [];
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "COMPANY"
    let sData = "";
    let myModal_newentry = "";
    let flag_btnsave = true;
    let flag_id_field = false;
    let flag_idcurr_field = false;
    let idcurr_field = "";
    let name_field = "";
    let owner_field = "";
    let email_field = "";
    let phone1_field = "";
    let phone2_field = "";
    let url1_field = "";
    let url2_field = "";
    let minimalfee_field = 0;
    let status_field = "";
    let create_field = "";
    let update_field = "";
    let idrecord = "";
    let pagingnow = 0;
    let searchHome = "";
    let filterHome = [];
    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_name
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.home_id
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
    }
    
    const NewData = (e,id,name,status,create,update) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            flag_id_field = true;
            flag_idcurr_field = false;
            idrecord = id;
            name_field = name;
            status_field = status;
            create_field = create;
            update_field = update;
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
    };
   
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    async function handleSave() {
        let flag = true
        let msg = ""
        
        if(sData == "New"){
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(idrecord == ""){
                flag = false
                msg += "The Code is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/companysave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"COMPANY-SAVE",
                    company_search: searchHome,
                    company_page: parseInt(pagingnow),
                    company_id: idrecord,
                    company_idcurr: idcurr_field,
                    company_name: name_field,
                    company_owner: owner_field,
                    company_phone1: phone1_field,
                    company_phone2: phone2_field,
                    company_email: email_field,
                    company_minfee: parseFloat(minimalfee_field),
                    company_url1: url1_field,
                    company_url2: url2_field,
                    company_status: status_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sData=="New"){
                    clearField()
                }
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    function clearField(){
        flag_id_field = false;
        idrecord = "";
        idcurr_field = "";
        name_field = "";
        owner_field = "";
        email_field = "";
        phone1_field = "";
        phone2_field = "";
        url1_field = "";
        url2_field = "";
        minimalfee_field = 0;
        status_field = "";
        create_field = "";
        update_field = "";
    }
  
    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData("New","","","");
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "SAVE":
                handleSubmit();break;
        }
    }
    const handleKeyboard_checkenter = (e) => {
        let keyCode = e.which || e.keyCode;
        if (keyCode === 13) {
                filterHome = [];
                listHome = [];
                const searchdata = {
                  searchHome,
                };
                dispatch("handleSearch", searchdata);
        }  
    };
    const handleSelectPaging = (event) => {
      let page = event.target.value;
      pagingnow = page;
      const pattern = {
        page
      };
      dispatch("handlePaging", pattern);
    };
    function status(e){
        let result = "DEACTIVE"
        if(e == "Y"){
            result = "ACTIVE"
        }
        return result
    }
</script>
<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-12">
            <Button
                on:click={callFunction}
                button_function="NEW"
                button_title="<i class='bi bi-plus-lg'></i>&nbsp;New"
                button_css="btn-dark"/>
            <Button
                on:click={callFunction}
                button_function="REFRESH"
                button_title="<i class='bi bi-arrow-clockwise'></i>&nbsp;Refresh"
                button_css="btn-primary"/>
            <Panel
                card_search={true}
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-title">
                    <div class="float-end">
                      <select
                        on:change={handleSelectPaging}
                        style="text-align: center;"
                        class="form-control">
                        {#each listPage as rec}
                          <option value={rec.page_value}>{rec.page_display}</option>
                        {/each}
                      </select>
                    </div>
                  </slot:template>
                <slot:template slot="card-search">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value={searchHome}
                            on:keypress={handleKeyboard_checkenter}
                            type="text"
                            class="form-control"
                            placeholder="Search Code, Departement"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" >&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="2%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">&nbsp;</th>
                                <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CODE</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DEPARTEMENT</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterHome as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                            //e,id,name,status,create,update
                                                NewData("Edit",rec.home_id, rec.home_name,rec.home_status,
                                                rec.home_create, rec.home_update);
                                            }} class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP  style="text-align: center;vertical-align: top;font-size: 11px;">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                                            {status(rec.home_status)}
                                        </span>
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_id}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_name}</td>
                                </tr>
                            {/each}
                        </tbody>
                        {:else}
                        <tbody>
                            <tr>
                                <td colspan="20">
                                    <center>
                                        <Loader />
                                    </center>
                                </td>
                            </tr>
                        </tbody>
                        {/if} 
                    </table>
                </slot:template>
            </Panel>
        </div>
    </div>
</div>

<Modal
	modal_id="modalentrycrud"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-md-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Currency</label>
                    <select
                        bind:value="{idcurr_field}" 
                        name="currency" id="currency" 
                        disabled={flag_idcurr_field} 
                        class="required form-control ">
                        <option value="">--Please Select--</option>
                        {#each listCurrency as rec}
                        <option value="{rec.curr_id}">{rec.curr_id}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">CODE</label>
                    <Input_custom
                        bind:value={idrecord}
                        input_tipe="text_uppercase_trim"
                        input_required="required"
                        input_maxlength="5"
                        disabled={flag_id_field}
                        input_placeholder="CODE"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Name</label>
                    <Input_custom
                        bind:value={name_field}
                        input_tipe="text_standart"
                        input_required="required"
                        input_maxlength="50"
                        input_placeholder="Name"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Owner</label>
                    <Input_custom
                        bind:value={owner_field}
                        input_tipe="text_standart"
                        input_required="required"
                        input_maxlength="50"
                        input_placeholder="Name"/>
                </div>
            </div>
            <div class="col-md-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Email</label>
                    <Input_custom
                        bind:value={email_field}
                        input_tipe="text_standart"
                        input_required="required"
                        input_maxlength="250"
                        input_placeholder="Email"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone1</label>
                    <Input_custom
                        bind:value={phone1_field}
                        input_tipe="text_standart"
                        input_required="required"
                        input_maxlength="20"
                        input_placeholder="Phone1"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone2</label>
                    <Input_custom
                        bind:value={phone2_field}
                        input_tipe="text_standart"
                        input_required=""
                        input_maxlength="20"
                        input_placeholder="Phone2"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Url 1</label>
                    <Input_custom
                        bind:value={url1_field}
                        input_tipe="text_standart"
                        input_required="required"
                        input_maxlength="250"
                        input_placeholder="Url1"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Url 2</label>
                    <Input_custom
                        bind:value={url2_field}
                        input_tipe="text_standart"
                        input_required=""
                        input_maxlength="250"
                        input_placeholder="url2"/>
                </div>
                <div class="mb-0">
                    <label for="exampleForm" class="form-label">Minimal Fee</label>
                    <Input_custom
                        bind:value={minimalfee_field}
                        input_tipe="number_float"
                        input_required=""
                        input_maxlength="5"
                        input_placeholder="Minimal Fee"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select
                        class="form-control required"
                        bind:value={status_field}>
                        <option value="">--Please Select--</option>
                        <option value="Y">ACTIVE</option>
                        <option value="N">DEACTIVE</option>
                    </select>
                </div>
                {#if sData != "New"}
                <div class="mb-3">
                    <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                        Create : {create_field}<br />
                        Update : {update_field}
                    </div>
                </div>
                {/if}
            </div>
        </div>
        
        
        
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSave();
            }} 
            button_function="SAVE"
            button_title="<i class='bi bi-save'></i>&nbsp;&nbsp;Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

