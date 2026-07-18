import {
    createSignal,
    For,
    Show,
} from "solid-js";

export interface SelectOption {
    value: string;
    label: string;
}

interface SelectProps {
    value: string;
    options: SelectOption[];
    onChange: (value: string) => void;
}

export default function Select(
    props: SelectProps
) {
    const [open, setOpen] =
        createSignal(false);

    const selectedOption = () =>
        props.options.find(
            (option) =>
                option.value === props.value
        );

    return (
        <div class="select">

            <button
                class="select__trigger"
                type="button"
                onClick={() =>
                    setOpen(!open())
                }
            >
                <span>
                    {selectedOption()?.label}
                </span>

                <span class="select__arrow">
                    ▾
                </span>
            </button>


            <Show when={open()}>

                <div class="select__menu">

                    <For each={props.options}>

                        {(option) => (

                            <button
                                class="select__option"
                                type="button"
                                classList={{
                                    selected:
                                        option.value ===
                                        props.value
                                }}
                                onClick={() => {

                                    props.onChange(
                                        option.value
                                    );

                                    setOpen(false);
                                }}
                            >
                                {option.label}
                            </button>

                        )}

                    </For>

                </div>

            </Show>

        </div>
    );
}