use crate::semantic::types::{MonoType, SubstitutionMap, Tvar};

// A substitution defines a function that takes a monotype as input
// and returns a monotype as output. The output type is interpreted
// as being equivalent to the input type.
//
// Substitutions are idempotent. Given a substitution s and an input
// type x, we have s(s(x)) = s(x).
//
#[derive(Debug, PartialEq)]
pub struct Substitution(SubstitutionMap);

// Derive a substitution from a hash map.
impl From<SubstitutionMap> for Substitution {
    fn from(values: SubstitutionMap) -> Substitution {
        Substitution(values)
    }
}

// The `allow` attribute below is a side effect of the orphan impl rule as
// well as the implicit_hasher lint. For more info, see
// https://github.com/rust-lang/rfcs/issues/1856

// Derive a hash map from a substitution.
#[allow(clippy::implicit_hasher)]
impl From<Substitution> for SubstitutionMap {
    fn from(sub: Substitution) -> SubstitutionMap {
        sub.0
    }
}

impl Substitution {
    pub fn empty() -> Substitution {
        Substitution(SubstitutionMap::new())
    }

    pub fn apply(&self, tv: Tvar) -> MonoType {
        match self.0.get(&tv) {
            Some(t) => t.clone(),
            None => MonoType::Var(tv),
        }
    }

    pub fn merge(self, with: Substitution) -> Substitution {
        let applied: SubstitutionMap = self
            .0
            .into_iter()
            .map(|(k, v)| (k, v.apply(&with)))
            .collect();
        Substitution(applied.into_iter().chain(with.0.into_iter()).collect())
    }
}

// A type is substitutable if a substitution can be applied to it.
pub trait Substitutable {
    fn apply(self, sub: &Substitution) -> Self;
    fn free_vars(&self) -> Vec<Tvar>;
}
