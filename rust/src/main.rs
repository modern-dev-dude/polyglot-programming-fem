enum RSEnum {
    FOO(i32),
    FOO2(Option<i32>),
    BAR(String),
    BAZ(Vec<String>),
}

fn main() {
    let foo = RSEnum::FOO2(Some(5));

    if let RSEnum::FOO(value) = foo {
        print!("foo: {}", value);
    }

    if let RSEnum::FOO2(Some(value)) = foo {
        print!("foo2: {}", value);
    }
}
