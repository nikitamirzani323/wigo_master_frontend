<script>
    export let value;
    export let input_tipe = "text";
    export let disabled = false;
    export let input_maxlength = 4;
    export let input_required = "";
    export let input_placeholder = "";

    const handleKeyboard_upppercase = (e) => {
        let str = e.target.value;
        let newstr = str.replace(" ","");
        return e.target.value = newstr.toUpperCase();
	}
    const handleKeyboard_number_standart = (e) => {
        if (isNaN(parseInt(e.target.value))) {
            return e.target.value = "";
        }else{
            return e.target.value = parseInt(e.target.value);
        }
	}
</script>
{#if input_tipe == "text_standart"}
    <input 
        bind:value
        {disabled}
        class="form-control {input_required}"
        maxlength="{input_maxlength}"
        type="text"
        placeholder="{input_placeholder}"/>
{/if}
{#if input_tipe == "text_uppercase_trim"}
    <input 
        bind:value
        on:keyup={handleKeyboard_upppercase}
        {disabled}
        class="form-control {input_required}"
        maxlength="{input_maxlength}"
        type="text"
        placeholder="{input_placeholder}"/>
{/if}
{#if input_tipe == "number_standart"}
    <input 
        bind:value
        on:keyup={handleKeyboard_number_standart}
        {disabled}
        class="form-control {input_required}"
        maxlength="{input_maxlength}"
        style="text-align: right;"
        type="text"
        placeholder="{input_placeholder}"/>
    <div id="passwordHelpBlock" class="form-text" style="text-align: right;color:blue;">
        {new Intl.NumberFormat().format(value)}
    </div>
{/if}